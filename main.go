package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)

	http.ListenAndServe(":8080", nil)
}
