#!/bin/sh

set -e
set -x

GOVERSION=1.11
GOFLAGS='-i -v -a'
export GOOS=js
export GOARCH=wasm

go${GOVERSION} build $GOFLAGS -o dagger.wasm dagger.go
go${GOVERSION} build $GOFLAGS -o nothing.wasm nothing.go
