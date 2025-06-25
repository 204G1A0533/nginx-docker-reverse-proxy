package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Service 1!")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}	

func main() {
    http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler) 
    fmt.Println("Service 1 listening on port 8001...")
    http.ListenAndServe(":8001", nil)
}
