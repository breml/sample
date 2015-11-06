package sample

import (
	"testing"
)

func BenchmarkReciprocalUint64Rand(b *testing.B) {
	s, err := NewReciprocalUint64(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint64 must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkReciprocalUint64From(b *testing.B) {
	s, err := NewReciprocalUint64(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint64 must not error", err)
	}

	benchmarkFrom(b, s)
}
