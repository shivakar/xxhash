package xxhash

import "encoding/binary"

// rotl64 performs circular rotation on uint64
func rotl64(x uint64, r uint) uint64 {
	return ((x << r) | (x >> (64 - r)))
}

// bToU64 converts a byte array into uint64
func bToU64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

// bToU32 converts a byte array into uint32
func bToU32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}
