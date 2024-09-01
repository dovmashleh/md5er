package md5

import "encoding/binary"

var InitVector HashVector = HashVector{0x67452301, 0xefcdab89, 0x98badcfe, 0x10325476}

// AsByteArray returns [16]byte as LittleEndian
func AsByteArray(mes []byte) [16]byte {
	mes = appendLenAndPad(mes)
	vector := InitVector
	for i := 0; i < len(mes); i += 64 {
		chunk := Chunk{abcd: vector}
		copy(chunk.chunkData[:], mes[i:i+64])
		vector = chunk.Process()
	}
	result := [16]byte{}
	binary.Encode(result[:], binary.LittleEndian, vector)
	return result
}
func AsByteSlice(mes []byte) []byte {
	array := AsByteArray(mes)
	return array[:]
}
func appendLenAndPad(mes []byte) []byte {
	mes = append(mes, 128)
	length := uint64(len(mes)) * 8
	pad := (512 + (448 - length%512)) % 512
	tail := make([]byte, pad/8)
	lenLeft := make([]byte, 4)
	lenRight := make([]byte, 4)
	binary.Encode(lenLeft, binary.LittleEndian, uint32((length-8)%(1<<32)))
	binary.Encode(lenRight, binary.LittleEndian, uint32((length-8)/(1<<32)))
	tail = append(tail, lenLeft...)
	tail = append(tail, lenRight...)
	mes = append(mes, tail...)
	return mes
}
