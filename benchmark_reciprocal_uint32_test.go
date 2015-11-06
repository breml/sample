package sample

import (
	"testing"
)

func BenchmarkReciprocalUint32Rand(b *testing.B) {
	s, err := NewReciprocalUint32(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint32 must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkReciprocalUint32From(b *testing.B) {
	s, err := NewReciprocalUint32(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint32 must not error", err)
	}

	benchmarkFrom(b, s)
}
