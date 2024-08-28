package main

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	var artists []Artist
	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")

		return
	}
	// Get the artist ID from the query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	patren := `\d+$`
	re := regexp.MustCompile(patren)
	if !re.MatchString(idStr) {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	var artist Artist
	if err := fetchData("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &artist); err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if err := fetchData("https://groupietrackers.herokuapp.com/api/locations/"+idStr, &artist.Locations); err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if err := fetchData("https://groupietrackers.herokuapp.com/api/dates/"+idStr, &artist.Dates); err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if err := fetchData("https://groupietrackers.herokuapp.com/api/relation/"+idStr, &artist.Relations); err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if artist.Name == "" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.ExecuteTemplate(w, "profile.html", artist)
	if err != nil {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
}

func renderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("static/error.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
	err = tmpl.Execute(w, struct {
		S int
		M string
	}{
		S: status,
		M: message,
	})
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
}
