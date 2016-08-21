package xxhash

import (
	"encoding/binary"
	"encoding/hex"
	"hash"
)

var (
	xxhash64 *XXHash64
	_        hash.Hash64 = xxhash64
	_        hash.Hash   = xxhash64
)

// Constants
const (
	prime64_1 = 11400714785074694791
	prime64_2 = 14029467366897019727
	prime64_3 = 1609587929392839161
	prime64_4 = 9650029242287828579
	prime64_5 = 2870177450012600261
)

// XXHash64 implements the 64-bit variant of the xxHash algorithm.
type XXHash64 struct {
	seed           uint64
	v1, v2, v3, v4 uint64
	len            uint64
	mem            [32]byte
	memsize        uint32
}

// String returns the current value of the hash as a hexadecimal string
func (x *XXHash64) String() string {
	return hex.EncodeToString(x.Sum(nil))
}

// Uint64 returns the current value of the hash as an uint64
func (x *XXHash64) Uint64() uint64 {
	return binary.BigEndian.Uint64(x.Sum(nil))
}

/*
 * Implement hash.Hash64 interface
 */

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying has state
func (x *XXHash64) Sum(b []byte) []byte {
	v := x.Sum64()
	return append(b, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32),
		byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// Reset resets the Hash to its initial state.
func (x *XXHash64) Reset() {
	x.v1 = x.seed + prime64_1 + prime64_2
	x.v2 = x.seed + prime64_2
	x.v3 = x.seed // + 0000
	x.v4 = x.seed - prime64_1
	x.len = 0
	x.memsize = 0
	for i := range x.mem {
		x.mem[i] = 0
	}
}

// Size returns the number of bytes Sum will return.
func (x *XXHash64) Size() int {
	return 8
}

// BlockSize returns the hash's underlying block size.
func (x *XXHash64) BlockSize() int {
	return 32
}

// Sum64 returns the current hash state
func (x *XXHash64) Sum64() uint64 {
	var h uint64

	if x.len >= 32 {
		v1, v2, v3, v4 := x.v1, x.v2, x.v3, x.v4

		h = rotl64_1(v1) + rotl64_7(v2) + rotl64_12(v3) + rotl64_18(v4)

		v1 *= prime64_2
		v1 = rotl64_31(v1)
		v1 *= prime64_1
		h ^= v1
		h = h*prime64_1 + prime64_4

		v2 *= prime64_2
		v2 = rotl64_31(v2)
		v2 *= prime64_1
		h ^= v2
		h = h*prime64_1 + prime64_4

		v3 *= prime64_2
		v3 = rotl64_31(v3)
		v3 *= prime64_1
		h ^= v3
		h = h*prime64_1 + prime64_4

		v4 *= prime64_2
		v4 = rotl64_31(v4)
		v4 *= prime64_1
		h ^= v4
		h = h*prime64_1 + prime64_4
	} else {
		h = x.seed + prime64_5
	}

	h += x.len

	if x.memsize > 0 {
		in := x.mem[:x.memsize:x.memsize]
		for len(in) >= 8 {
			h ^= rotl64_31(bToU64(in[:8:8])*prime64_2) * prime64_1
			h = rotl64_27(h)*prime64_1 + prime64_4
			in = in[8:len(in):len(in)]
		}
		if len(in) >= 4 {
			h ^= uint64(bToU32(in[:4:4])) * prime64_1
			h = rotl64_23(h)*prime64_2 + prime64_3
			in = in[4:len(in):len(in)]
		}
		for i := 0; i < len(in); i++ {
			h ^= uint64(uint8(in[i])) * prime64_5
			h = rotl64_11(h) * prime64_1
		}
	}

	h ^= h >> 33
	h *= prime64_2
	h ^= h >> 29
	h *= prime64_3
	h ^= h >> 32

	return h
}

// Write adds more data to the running hash and updates the hash state
// It never returns an error
func (x *XXHash64) Write(input []byte) (int, error) {
	l := len(input)
	x.len += uint64(l)

	if x.memsize+uint32(l) < 32 {
		// new data fits into the buffer
		x.memsize += uint32(copy(x.mem[x.memsize:], input))
		return l, nil
	}

	if x.memsize > 0 {
		// new data does not fit into the buffer
		// and some data is still unprocessed from previous update
		n := 32 - x.memsize
		copy(x.mem[x.memsize:], input[:n:len(input)])

		x.v1 += bToU64(x.mem[:8:8]) * prime64_2
		x.v1 = rotl64_31(x.v1) * prime64_1

		x.v2 += bToU64(x.mem[8:16:16]) * prime64_2
		x.v2 = rotl64_31(x.v2) * prime64_1

		x.v3 += bToU64(x.mem[16:24:24]) * prime64_2
		x.v3 = rotl64_31(x.v3) * prime64_1

		x.v4 += bToU64(x.mem[24:32:32]) * prime64_2
		x.v4 = rotl64_31(x.v4) * prime64_1

		input = input[n:len(input):len(input)]
		x.memsize = 0
	}

	if len(input) >= 32 {
		for len(input) >= 32 {
			x.v1 += bToU64(input[:8:8]) * prime64_2
			x.v1 = rotl64_31(x.v1) * prime64_1

			x.v2 += bToU64(input[8:16:16]) * prime64_2
			x.v2 = rotl64_31(x.v2) * prime64_1

			x.v3 += bToU64(input[16:24:24]) * prime64_2
			x.v3 = rotl64_31(x.v3) * prime64_1

			x.v4 += bToU64(input[24:32:32]) * prime64_2
			x.v4 = rotl64_31(x.v4) * prime64_1

			input = input[32:len(input):len(input)]
		}
	}

	if len(input) > 0 {
		x.memsize += uint32(copy(x.mem[x.memsize:], input))
	}

	return l, nil
}

// NewSeedXXHash64 returns an instance of XXHash64 with the specified seed.
func NewSeedXXHash64(seed uint64) *XXHash64 {
	x := &XXHash64{
		seed: seed,
	}
	x.Reset()
	return x
}

// NewXXHash64 returns an instance of XXHash64 with seed set to 0.
func NewXXHash64() *XXHash64 {
	return NewSeedXXHash64(0)
}
