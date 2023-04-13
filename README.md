# WebZip 
WebZip is implemented in Golang and uses wasm technology to perform file compression in the browser. It does not upload the files to the server.

## How to use
- run `go run examples/main.go` 

- open `127.0.0.1:8080`

## How to build
- `GOARCH=wasm GOOS=js go build -o compress.wasm wasm/compress.go` 