// +build js,wasm ignore

package main

import (
	"strconv"

	"github.com/Xe/olin/dagger"
)

func main() {
	fd := dagger.OpenFile("fd://1")
	println(strconv.Itoa(int(fd)))
}
