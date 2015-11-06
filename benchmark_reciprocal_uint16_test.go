package sample

import (
	"testing"
)

func BenchmarkReciprocalUint16Rand(b *testing.B) {
	s, err := NewReciprocalUint16(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint16 must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkReciprocalUint16From(b *testing.B) {
	s, err := NewReciprocalUint16(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint16 must not error", err)
	}

	benchmarkFrom(b, s)
}
