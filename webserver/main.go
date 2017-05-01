package main

import (
  "net/http"
)

const (
  port = ":8080"
)

func Show(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World !"))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", Show)
  http.ListenAndServe(port, mux)
}
