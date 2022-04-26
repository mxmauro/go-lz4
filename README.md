go-lz4
======

go-lz4 is Go minimal wrapper for the [LZ4](https://github.com/lz4/lz4) compression library.

## Usage with example

```golang
package example

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"github.com/randlabs/go-lz4"
)

func main() {
	var outputSize int

	// Generate random input data
	data := make([]byte, 4096)
	rand.Read(data)

	// Compress the data. You can also call to "MaxCompressedSize" and preallocate a byte buffer with the returned size 
	compressed, err := lz4.AllocAndCompress(data)
	if err != nil {
		fmt.Printf("Unable to allocate and compress random data [err=%v]\n", err.Error())
		return
	}

	// Now decompress the compressed data. We allocate the 4096 byte buffer because we know the size of the decompressed data
	decompressed := make([]byte, 4096)
	outputSize, err = lz4.Decompress(compressed, decompressed)
	if err != nil {
		fmt.Printf("Unable to decompress the compressed data [err=%v]\n", err.Error())
		return
	}

	// At last, check if the decompressed data and the original one are equal
	if !bytes.Equal(data, decompressed[:outputSize]) {
		fmt.Printf("The original and decompressed data does not match\n")
		return
	}
}

```

## CGO
The library uses the original C native implementation. To use it, CGO must be enabled.
Also take into account cross-platform compilation may fail if `gcc` utility for the target OS is not installed.

## License
See `LICENSE` file for details.
