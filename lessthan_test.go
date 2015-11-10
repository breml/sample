package sample

import (
	"fmt"
	"math"
	"testing"
)

func TestLessThanNilSampler(t *testing.T) {
	var sampler Sampler
	var err error

	sampler, err = NewLessThan(0)
	if err == nil {
		t.Fatal("NewLessThan must error if rate = 0", err)
	}
	_ = sampler
}

func TestLessThanRate1(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1
	sampler, err = NewLessThan(1)
	if err != nil {
		t.Fatal("NewLessThan must not error", err)
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

func TestLessThanSampleFrom(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1000 with from
	sampler, err = NewLessThan(1000)
	if err != nil {
		t.Fatal("NewLessThan must not error", err)
	}
	// true
	if sampler.SampleFrom(0) != true {
		t.Error("sampling 0 with rate 1000 did not return true")
	}
	// false
	if sampler.SampleFrom(math.MaxUint64) != false {
		t.Error("sampling MaxUint64 with rate 1000 did not return false")
	}
}

func TestLessThanSeedSample(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample with seed, false
	sampler, err = NewLessThanSeeded(1000, 0)
	if err != nil {
		t.Fatal("NewLessThan must not error", err)
	}
	if sampler.Sample() != false {
		t.Error("sampling with seed 0 and rate 1000 did not return false")
	}

	// Sample with seed, true
	sampler, err = NewLessThanSeeded(1000, 165)
	if err != nil {
		t.Fatal("NewLessThan must not error", err)
	}
	if sampler.Sample() != true {
		t.Error("sampling with seed 165 and rate 1000 did not return true")
	}
}

func TestLessThanString(t *testing.T) {
	var sampler Sampler

	sampler = &lessThan{sampleState: sampleState{rate: 10, seed: 0, sampleCount: 1000, trueCount: 100}, boundary: math.MaxUint64 / 10}

	if sampler.String() != "&{sampleState:{rate:10 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>} boundary:1844674407370955161}" {
		t.Error("Expected: &{sampleState:{rate:10 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>} boundary:1844674407370955161}, Got:", sampler.String())
	}
}

func ExampleNewLessThanSeeded() {
	s, err := NewLessThanSeeded(10, 0)
	if err != nil {
		fmt.Println("Unable to initialize sampler", err)
	}
	for i := 0; i < 100; i++ {
		if s.Sample() {
			fmt.Println(i, "got sampled by LessThan sampler")
		}
	}
	fmt.Println(Stats(s))
	// Output:
	// 25 got sampled by LessThan sampler
	// 27 got sampled by LessThan sampler
	// 29 got sampled by LessThan sampler
	// 43 got sampled by LessThan sampler
	// 54 got sampled by LessThan sampler
	// 73 got sampled by LessThan sampler
	// Rate: 10, SampleCount: 100, TrueCount: 6, Deviation: -66.6667%
}
