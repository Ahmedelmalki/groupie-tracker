package main

import (
	"net/http"
	"text/template"
)

type ArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}
	var artists []Artist
	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", struct {
		Artists []Artist
	}{
		Artists: artists,
	})
	if err != nil {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}
	// Get the artist ID from the query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}
	var artist Artist
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &artist)
	if err != nil {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}

	err = tmpl.ExecuteTemplate(w, "profile.html", artist)
	if err != nil {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}
}

func babyError(w http.ResponseWriter, statusCode int) {
	tmpl, err := template.ParseFiles("static/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		StatusCode int
		Message    string
	}{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode),
	}

	w.WriteHeader(statusCode)
	err = tmpl.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
