package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//handle routes
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/profile", handleProfile)
	//start the server
	fmt.Println("Server is running at http://localhost:8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
