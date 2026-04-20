package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", HelloServer)
  http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<html><body><h1>Hello World!</h1></body></html>")
}

// run: 
// go build app.go
// then:
// ./app
// visit:
// http://localhost/