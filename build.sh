#!/bin/bash
GOOS=js GOARCH=wasm go build -o test.wasm class.go
GOOS=js GOARCH=wasm go build -o wrapper.wasm wrapper.go test_struct.go
curl -sO https://raw.githubusercontent.com/golang/go/go1.12.7/misc/wasm/wasm_exec.js
