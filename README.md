how to access Go's struct from wasm

# how to use
## build
```
./build.sh
```

## execute
```
go run server.go
```

One exasmple in http://localhost:8080/wasm_exec.html shows how to use Go struct from WASM.


Another example is available in http://localhost:8080/wrapper.html 
This example shows how to use Go's struct without any modification. I implemented wrapper function.

It also shows how to pass struct declared in Go as an argument. 

