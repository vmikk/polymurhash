// Package polymurhash implements PolymurHash v2, a fast keyed 64-bit
// almost-universal hash designed for hash tables and related data structures.
//
// A Seed identifies a hash function and may be reused concurrently. Use
// MakeSeed when constructing a keyed data structure, or a deterministic seed
// constructor when hashes must be reproducible.
//
// PolymurHash is not a cryptographic hash and must not be used where attackers
// can observe hash outputs or request hashes of chosen inputs.
//
// This package is an altered Go implementation of PolymurHash v2 by Orson
// Peters. See LICENSE-POLYMUR-HASH for the upstream zlib license notice.
package polymurhash
