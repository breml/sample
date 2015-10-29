package sample

import (
	"testing"
)

func BenchmarkLessThanRand(b *testing.B) {
	s, err := NewLessThan(1000)
	if err != nil {
		b.Fatal("NewLessThan must not error", err)
	}

	// run sample b.N times
	for n := 0; n < b.N; n++ {
		s.Sample()
	}

	b.Log(Stats(s))
}

func BenchmarkLessThanFrom(b *testing.B) {
	s, err := NewLessThan(1000)
	if err != nil {
		b.Fatal("NewLessThan must not error", err)
	}

	// run sample b.N times
	for n := 0; n < b.N; n++ {
		s.SampleFrom(uint64(n))
	}
}
