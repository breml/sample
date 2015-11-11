package sample

import (
	"fmt"
	"math"
	"testing"
)

func TestModuloNilSampler(t *testing.T) {
	var sampler Sampler
	var err error

	sampler, err = NewModulo(0)
	if err == nil {
		t.Fatal("NewModulo must error if rate = 0", err)
	}

	_ = sampler
}

func TestModuloRate1(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1
	sampler, err = NewModulo(1)
	if err != nil {
		t.Fatal("NewModulo must not error", err)
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

func TestModuloSampleFrom(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample 1024 with from
	sampler, err = NewModulo(1024)
	if err != nil {
		t.Fatal("NewModulo must not error", err)
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

func TestModuloSeedSample(t *testing.T) {
	var sampler Sampler
	var err error

	// Sample with seed, false
	sampler, err = NewModuloSeeded(1024, 0)
	if err != nil {
		t.Fatal("NewModulo must not error", err)
	}
	if sampler.Sample() != false {
		t.Error("sampling with seed 0 and rate 1024 did not return false")
	}

	// Sample with seed, true
	sampler, err = NewModuloSeeded(1024, 643)
	if err != nil {
		t.Fatal("NewModulo must not error", err)
	}

	if sampler.Sample() != true {
		t.Error("sampling with seed 643 and rate 1024 did not return true")
	}
}

func TestModuloString(t *testing.T) {
	var sampler Sampler

	sampler = &modulo{sampleState: sampleState{rate: 1024, seed: 0, sampleCount: 1000, trueCount: 100}}

	if sampler.String() != "&{sampleState:{rate:1024 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>}}" {
		t.Error("Expected: &{sampleState:{rate:1024 seed:0 sampleCount:1000 trueCount:100 rnd:<nil>}}, Got:", sampler.String())
	}
}

func ExampleNewModuloSeeded() {
	s, err := NewModuloSeeded(10, 0)
	if err != nil {
		fmt.Println("Unable to initialize sampler", err)
	}
	for i := 0; i < 100; i++ {
		if s.Sample() {
			fmt.Println(i, "got sampled by Modulo sampler")
		}
	}
	fmt.Println(Stats(s))
	// Output:
	// 12 got sampled by Modulo sampler
	// 31 got sampled by Modulo sampler
	// 34 got sampled by Modulo sampler
	// 60 got sampled by Modulo sampler
	// Rate: 10, SampleCount: 100, TrueCount: 4, Deviation: -150.0000%

}
