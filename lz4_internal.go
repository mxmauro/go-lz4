// See the LICENSE file for license details.

package lz4

import (
	"unsafe"

	"C"
)

// -----------------------------------------------------------------------------

func addressOfBytes(data []byte) *C.char {
	if len(data) > 0 {
		return (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		return (*C.char)(unsafe.Pointer(nil))
	}
}
