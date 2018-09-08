// Package abi is the common interface for all ABI implementations.
package abi

import (
	"github.com/Xe/olin/internal/abi/file"
	"github.com/perlin-network/life/exec"
)

type Namer = file.Namer

// ABI is the interface that is provided to a webaseembly module. Instances
// of this interface will live for the lifetime of webassembly modules.
//
// This is low level intentionally.
type ABI interface {
	exec.ImportResolver
	Namer

	Files() []File
	Open(File)
}

type File = file.File
