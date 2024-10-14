// See the LICENSE file for license details.

package lz4_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/mxmauro/go-lz4"
)

// -----------------------------------------------------------------------------

func TestCompress(t *testing.T) {
	var outputSize int

	data := make([]byte, 4096)
	rand.Read(data)

	compressed, err := lz4.AllocAndCompress(data)
	if err != nil {
		t.Fatalf("unable to allocate and compress [err=%v]", err.Error())
	}

	decompressed := make([]byte, 4096)
	outputSize, err = lz4.Decompress(compressed, decompressed)
	if err != nil {
		t.Fatalf("unable to decompress [err=%v]", err.Error())
	}

	if !bytes.Equal(data, decompressed[:outputSize]) {
		t.Fatal("original and decompressed data mismatch")
	}
}

func TestCompressWithSmallBuffer(t *testing.T) {
	data := make([]byte, 4096)
	rand.Read(data)

	compressed := make([]byte, 1)
	_, err := lz4.Compress(data, compressed)
	if err == nil {
		t.Fatalf("buffer too small error expected")
	}
}
