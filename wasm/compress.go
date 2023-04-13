//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"syscall/js"
)

//export MyCompress
func MyCompress(this js.Value, inputs []js.Value) interface{} {

	defer func() {
		fmt.Println("go compress done!")
	}()

	// 将js.Value转换为[]byte
	dataTypedArray := js.Global().Get("Uint8Array").New(inputs[0])
	//defer dataTypedArray.Release()

	data := make([]byte, dataTypedArray.Length())
	js.CopyBytesToGo(data, dataTypedArray)

	// 创建一个字节数组缓冲区
	var buf bytes.Buffer
	// 将压缩器写入缓冲区
	gz := gzip.NewWriter(&buf)

	if _, err := gz.Write(data); err != nil {
		return nil
	}
	if err := gz.Flush(); err != nil {
		fmt.Println("flush err:", err)
		return nil
	}
	// 这里如果用 defer 会失效
	if err := gz.Close(); err != nil {
		fmt.Println("close err:", err)
		return nil
	}

	// 获取压缩后的数据
	compressedData := buf.Bytes()

	compressedDataValue := js.Global().Get("Uint8Array").New(len(compressedData))
	js.CopyBytesToJS(compressedDataValue, compressedData)
	return compressedDataValue

}

func main() {
	js.Global().Set("MyCompress", js.FuncOf(MyCompress))

	<-make(chan bool)
}

// GOARCH=wasm GOOS=js go build -o compress.wasm wasm/compress.go
