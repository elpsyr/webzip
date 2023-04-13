package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取当前目录
		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 读取index.html文件
		body, err := ioutil.ReadFile(dir + "/examples/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 将文件内容写入响应
		fmt.Fprint(w, string(body))
	})

	http.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		// 获取当前目录
		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 读取wasm_exec.js文件
		body, err := ioutil.ReadFile(dir + "/examples/wasm_exec.js")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 将文件内容写入响应
		fmt.Fprint(w, string(body))
	})

	http.HandleFunc("/compress.wasm", func(w http.ResponseWriter, r *http.Request) {
		// 获取当前目录
		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 读取compress.wasm文件
		body, err := ioutil.ReadFile(dir + "/examples/compress.wasm")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 将文件内容写入响应
		fmt.Fprint(w, string(body))
	})

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
