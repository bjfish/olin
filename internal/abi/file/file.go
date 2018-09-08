// Package file exists to hack around stuff.
package file

// Namer is something that has a name.
type Namer interface {
	Name() string
}

type Closer interface {
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

// File is the most common denominator for most of what you usefully need out of
// files to make useful programs.
type File interface {
	Closer
	Reader
	Writer
	Namer

	// Sync isn't required for all backends, but it is used for some, such as HTTP.
	Sync() error
}
