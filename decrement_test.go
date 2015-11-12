package sample

import (
	"fmt"
	"math"
	"testing"
)

func TestDecrementNilSampler(t *testing.T) {
	var sampler Sampler
	var err error

	sampler, err = NewDecrement(0)
	if err == nil {
		t.Fatal("NewDecrement must error if rate = 0", err)
	}
	_ = sampler
}

func TestDecrementRate1(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1
	sampler, err = NewDecrement(1)
	if err != nil {
		t.Fatal("NewDecrement must not error", err)
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

func TestDecrementSampleFrom(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1000 with from
	sampler, err = NewDecrement(1000)
	if err != nil {
		t.Fatal("NewDecrement must not error", err)
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

func TestDecrementSeedSample(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample with seed, false
	sampler, err = NewDecrementSeeded(1000, 0)
	if err != nil {
		t.Fatal("NewDecrement must not error", err)
	}
	if sampler.Sample() != false {
		t.Error("sampling with seed 0 and rate 1000 did not return false")
	}

	// Sample with seed, true
	sampler, err = NewDecrementSeeded(1000, 3859)
	if err != nil {
		t.Fatal("NewDecrement must not error", err)
	}
	if sampler.Sample() != true {
		t.Error("sampling with seed 3859 and rate 1000 did not return true")
	}
}

func TestDecrementString(t *testing.T) {
	var sampler Sampler

	sampler = &decrement{sampleState: sampleState{rate: 10, seed: 0, sampleCount: 1000, trueCount: 100}, count: 0}

	if sampler.String() != "&{sampleState:{rate:10 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>} count:0}" {
		t.Error("Expected: &{sampleState:{rate:10 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>} count:0}, Got:", sampler.String())
	}
}

func ExampleNewDecrementSeeded() {
	s, err := NewDecrementSeeded(10, 0)
	if err != nil {
		fmt.Println("Unable to initialize sampler", err)
	}
	for i := 0; i < 100; i++ {
		if s.Sample() {
			fmt.Println(i, "got sampled by Decrement sampler")
		}
	}
	fmt.Println(Stats(s))
	// Output:
	// 5 got sampled by Decrement sampler
	// 23 got sampled by Decrement sampler
	// 31 got sampled by Decrement sampler
	// 37 got sampled by Decrement sampler
	// 45 got sampled by Decrement sampler
	// 63 got sampled by Decrement sampler
	// 68 got sampled by Decrement sampler
	// 79 got sampled by Decrement sampler
	// 87 got sampled by Decrement sampler
	// Rate: 10, SampleCount: 100, TrueCount: 9, Deviation: -11.1111%

}
