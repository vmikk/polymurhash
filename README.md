# polymurhash

Package `polymurhash` implements [PolymurHash v2](https://github.com/orlp/polymur-hash), a fast keyed 64-bit almost-universal hash designed for hash tables and related data structures.

This is a Go port of the reference algorithm. Its outputs are tested against the upstream C vectors and additional binary boundary vectors generated from the reference implementation.

## Usage

Create one random seed for each independently keyed data structure, then reuse it for hashing. A `Seed` is immutable and safe to reuse concurrently.

```go
seed := polymurhash.MakeSeed()

const tableTweak = uint64(1)
bucketHash := polymurhash.String(seed, "customer:42", tableTweak)
_ = bucketHash
```

Use deterministic constructors when output must be reproducible, such as in tests or persistent formats:

```go
seed128 := polymurhash.NewSeed(0x0123456789abcdef, 0xfedcba9876543210)
seed64 := polymurhash.NewSeedFromUint64(0x0123456789abcdef)

sum1 := polymurhash.Bytes(seed128, []byte("key"), 0)
sum2 := polymurhash.String(seed64, "key", 0)
_, _ = sum1, sum2
```

The per-call `tweak` cheaply differentiates uses of one seed, for example separate hash tables. The reference algorithm does not make collision guarantees between different tweaks.

PolymurHash is not a cryptographic hash. Do not use it where an attacker can observe hash values or submit chosen inputs and rely on cryptographic resistance.

## API

```go
type Seed struct { /* unexported parameters */ }

func MakeSeed() Seed
func NewSeed(kSeed, sSeed uint64) Seed
func NewSeedFromUint64(seed uint64) Seed

func Bytes(seed Seed, b []byte, tweak uint64) uint64
func String(seed Seed, s string, tweak uint64) uint64
```

The zero value of `Seed` is invalid; construct seeds using one of the provided functions.

## License

The Go implementation is licensed under the MIT License in [LICENSE](LICENSE). The initial Go implementation was written by [Yonashiro Nao](https://github.com/orlp/polymur-hash). It is derived from Orson Peters' PolymurHash v2 reference implementation,licensed under the zlib license reproduced in [LICENSE-POLYMUR-HASH](LICENSE-POLYMUR-HASH).
