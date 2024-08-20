package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/profile", handleIndex)

	http.ListenAndServe(":8000", nil)
}
