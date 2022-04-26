package lz4

// #cgo CFLAGS: -O3 -Wall
// #include "lz4.h"
import "C"

import (
	"errors"
)

// -----------------------------------------------------------------------------

// MaxCompressedSize returns the maximum size that LZ4 compression may output.
func MaxCompressedSize(inputSize int) int {
	return int(C.LZ4_compressBound(C.int(inputSize)))
}

// AllocAndCompress compresses a buffer and returns the compressed data or an error.
func AllocAndCompress(source []byte) ([]byte, error) {
	dest := make([]byte, MaxCompressedSize(len(source)))
	n, err := Compress(source, dest)
	if err != nil {
		return nil, err
	}
	return dest[:n], nil
}

// Compress compresses a buffer into pre-allocated destination buffer. It returns the number of bytes
// written or an error.
func Compress(source []byte, dest []byte) (int, error) {
	n := int(C.LZ4_compress_default(addressOfBytes(source), addressOfBytes(dest), C.int(len(source)), C.int(len(dest))))
	if n <= 0 {
		return 0, errors.New("buffer too small")
	}
	return n, nil
}

// Decompress decompresses a buffer into a pre-allocated destination buffer. It returns the number of bytes
// written or an error.
func Decompress(source []byte, dest []byte) (int, error) {
	n := int(C.LZ4_decompress_safe(addressOfBytes(source), addressOfBytes(dest), C.int(len(source)), C.int(len(dest))))
	if n < 0 {
		return 0, errors.New("buffer too small or malformed LZ4")
	}
	return n, nil
}
