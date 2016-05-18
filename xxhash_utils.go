package xxhash

import "encoding/binary"

// rotl64 performs circular rotation on uint64
//func rotl64(x uint64, r uint) uint64 {
//	return ((x << r) | (x >> (64 - r)))
//}

// Forcing the go compiler to generate _rotl instructions
func rotl64_1(x uint64) uint64 {
	return ((x << 1) | (x >> (64 - 1)))
}
func rotl64_7(x uint64) uint64 {
	return ((x << 7) | (x >> (64 - 7)))
}
func rotl64_11(x uint64) uint64 {
	return ((x << 11) | (x >> (64 - 11)))
}
func rotl64_12(x uint64) uint64 {
	return ((x << 12) | (x >> (64 - 12)))
}
func rotl64_18(x uint64) uint64 {
	return ((x << 18) | (x >> (64 - 18)))
}
func rotl64_23(x uint64) uint64 {
	return ((x << 23) | (x >> (64 - 23)))
}
func rotl64_27(x uint64) uint64 {
	return ((x << 27) | (x >> (64 - 27)))
}
func rotl64_31(x uint64) uint64 {
	return ((x << 31) | (x >> (64 - 31)))
}

// bToU64 converts a byte array into uint64
func bToU64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

// bToU32 converts a byte array into uint32
func bToU32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}
