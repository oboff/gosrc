// main
package main

import (
	"io"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<html><h1>hello world!\n</h1></html>")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":8080", nil) //http.FileServer(http.Dir(".")))
}
