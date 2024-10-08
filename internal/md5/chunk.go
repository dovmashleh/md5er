package md5

import (
	"encoding/binary"
	"math/bits"
)

var (
	s = [64]int{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}

	k = [64]uint32{
		0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
		0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
		0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
		0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
		0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
		0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
		0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
		0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
		0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
	}
)

type HashVector [4]uint32

func (vec HashVector) toABCD() (uint32, uint32, uint32, uint32) {
	return vec[0], vec[1], vec[2], vec[3]
}
func (vec *HashVector) plus(vecPlus [4]uint32) HashVector {
	for i := 0; i < 4; i++ {
		vec[i] += vecPlus[i]
	}
	return *vec
}

type Chunk struct {
	chunkData [64]byte
	abcd      HashVector
}

func (chunk Chunk) Process() HashVector {
	M := [16]uint32{}
	binary.Decode(chunk.chunkData[:], binary.LittleEndian, &M) //spliting chunk into 16 uint32 mes
	A, B, C, D := chunk.abcd.toABCD()
	for i := 0; i < 64; i++ {
		var F uint32
		var g int
		if i < 16 {
			F = (B & C) | (bNot(B) & D)
			g = i
		} else if i < 32 {
			F = (D & B) | (bNot(D) & C)
			g = (i*5 + 1) % 16
		} else if i < 48 {
			F = B ^ C ^ D
			g = (3*i + 5) % 16
		} else {
			F = C ^ (B | bNot(D))
			g = (7 * i) % 16
		}
		F = F + A + k[i] + M[g]
		A = D
		D = C
		C = B
		B = B + bits.RotateLeft32(F, s[i])
	}
	return chunk.abcd.plus([4]uint32{A, B, C, D})
}

func bNot(val uint32) uint32 {
	return val ^ (1<<32 - 1)
}
