package polymurhash

import (
	"fmt"
	"testing"
)

var (
	benchmarkHashResult uint64
	benchmarkSeedResult Seed
)

func BenchmarkBytes(b *testing.B) {
	seed := NewSeedFromUint64(0xfedbca9876543210)
	const tweak = uint64(0xabcdef0123456789)
	for _, length := range []int{4, 8, 16, 49, 50, 100, 1024, 4096} {
		input := boundaryInput(length)
		b.Run(fmt.Sprintf("%dB", length), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(length))
			var result uint64
			for i := 0; i < b.N; i++ {
				result = Bytes(seed, input, tweak)
			}
			benchmarkHashResult = result
		})
	}
}

func BenchmarkString(b *testing.B) {
	seed := NewSeedFromUint64(0xfedbca9876543210)
	const tweak = uint64(0xabcdef0123456789)
	for _, length := range []int{4, 8, 16, 49, 50, 100, 1024, 4096} {
		input := string(boundaryInput(length))
		b.Run(fmt.Sprintf("%dB", length), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(length))
			var result uint64
			for i := 0; i < b.N; i++ {
				result = String(seed, input, tweak)
			}
			benchmarkHashResult = result
		})
	}
}

func BenchmarkNewSeed(b *testing.B) {
	b.ReportAllocs()
	var seed Seed
	for i := 0; i < b.N; i++ {
		seed = NewSeed(uint64(i), uint64(i)+1)
	}
	benchmarkSeedResult = seed
}

func BenchmarkNewSeedFromUint64(b *testing.B) {
	b.ReportAllocs()
	var seed Seed
	for i := 0; i < b.N; i++ {
		seed = NewSeedFromUint64(uint64(i))
	}
	benchmarkSeedResult = seed
}
