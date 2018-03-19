#!/usr/bin/env bash

cd /go/src/github.com/llofberg/opa-rust

go build -buildmode=c-archive -o libopa.a libopa.go && \
rustc main.rs -L . -lopa && \
./main
