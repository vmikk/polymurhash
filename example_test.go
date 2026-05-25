package polymurhash_test

import "github.com/vmikk/polymurhash"

func Example() {
	seed := polymurhash.MakeSeed()

	const tableTweak = uint64(1)
	_ = polymurhash.String(seed, "customer:42", tableTweak)
	_ = polymurhash.Bytes(seed, []byte("customer:42"), tableTweak)
}

func ExampleNewSeedFromUint64() {
	seed := polymurhash.NewSeedFromUint64(0x0123456789abcdef)
	_ = polymurhash.String(seed, "stable key", 0)
}
