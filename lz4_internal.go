package lz4

import (
	"C"
	"unsafe"
)

// -----------------------------------------------------------------------------

func addressOfBytes(data []byte) *C.char {
	if len(data) > 0 {
		return (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		return (*C.char)(unsafe.Pointer(nil))
	}
}
