package sample

import (
	"testing"
)

func benchmarkRand(b *testing.B, s Sampler) {
	// run sample b.N times
	for n := 0; n < b.N; n++ {
		s.Sample()
	}

	b.Log(Stats(s))
}

func benchmarkFrom(b *testing.B, s Sampler) {
	// run sample b.N times
	for n := 0; n < b.N; n++ {
		s.SampleFrom(uint64(n))
	}
}

func BenchmarkLessThanRand(b *testing.B) {
	s, err := NewLessThan(1000)
	if err != nil {
		b.Fatal("NewLessThan must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkLessThanFrom(b *testing.B) {
	s, err := NewLessThan(1000)
	if err != nil {
		b.Fatal("NewLessThan must not error", err)
	}

	benchmarkFrom(b, s)
}

func BenchmarkPowerOf2Rand(b *testing.B) {
	s, err := NewPowerOf2(1024)
	if err != nil {
		b.Fatal("NewPowerOf2 must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkPowerOf2From(b *testing.B) {
	s, err := NewPowerOf2(1024)
	if err != nil {
		b.Fatal("NewPowerOf2 must not error", err)
	}

	benchmarkFrom(b, s)
}

func BenchmarkModuloRand(b *testing.B) {
	s, err := NewModulo(1000)
	if err != nil {
		b.Fatal("NewModulo must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkModuloFrom(b *testing.B) {
	s, err := NewModulo(1000)
	if err != nil {
		b.Fatal("NewModulo must not error", err)
	}

	benchmarkFrom(b, s)
}

func BenchmarkDecrementRand(b *testing.B) {
	s, err := NewDecrement(1000)
	if err != nil {
		b.Fatal("NewDecrement must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkDecrementFrom(b *testing.B) {
	s, err := NewDecrement(1000)
	if err != nil {
		b.Fatal("NewDecrement must not error", err)
	}

	benchmarkFrom(b, s)
}
