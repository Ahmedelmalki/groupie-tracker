package main

import (
	"net/http"
	"text/template"
)

type ArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", struct {
		Artists []Artist
	}{
		Artists: artists,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	// Get the artist ID from the query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Artist ID not provided", http.StatusBadRequest)
		return
	}
	var artist Artist
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &artist)
	if err != nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "profile.html", artist)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
