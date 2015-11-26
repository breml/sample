//+build generate
//go:generate ./gen_reciprocals.sh 8
//go:generate ./gen_reciprocals.sh 16
//go:generate ./gen_reciprocals.sh 32
//go:generate ./gen_reciprocals.sh 64

package sample

import (
	"testing"
)

func BenchmarkReciprocalUint8Rand(b *testing.B) {
	s, err := NewReciprocalUint8(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint8 must not error", err)
	}

	benchmarkRand(b, s)
}

func BenchmarkReciprocalUint8From(b *testing.B) {
	s, err := NewReciprocalUint8(1000)
	if err != nil {
		b.Fatal("NewReciprocalUint8 must not error", err)
	}

	benchmarkFrom(b, s)
}
