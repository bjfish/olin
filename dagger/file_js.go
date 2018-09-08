// +build js,wasm go1.11

package dagger

import (
	"errors"
	"strconv"

	"github.com/Xe/olin/internal/abi/file"
)

func OpenFile(furl string) int64 {
	return openFD(furl)
}

func openFD(furl string) int64

func Open(furl string) (file.File, error) {
	fd := openFD(furl)
	if fd < 0 {
		return nil, Error{Errno: Errno(fd)}
	}

	return File{fd: fd}, nil
}

// File is an externally managed file.
type File struct {
	file.File
	fd int64
}

func read(fd int64, buf []byte) int

func (f File) Read(buf []byte) (int, error) {
	n := read(f.fd, buf)
	if n < 0 {
		return n, errors.New("dagger: error code " + strconv.Itoa(n*-1))
	}

	return n, nil
}
