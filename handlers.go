package main

import (
	"fmt"
	"net/http"
	"regexp"
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
	patren := `\d+$`
	re := regexp.MustCompile(patren)
	if !re.MatchString(idStr) {
		babyError(w, http.StatusNotFound)
		return
	}
	if idStr == "" {
		babyError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(idStr)
	var artist Artist
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &artist)
	if err != nil {
		babyError(w, http.StatusInternalServerError)
		return
	}
	if artist.Id == 0 {
		http.Error(w, "This Goupe DOES NOT EXIST", 404)
		return
	}
	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		babyError(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "profile.html", artist)
	if err != nil {
		babyError(w, http.StatusNotFound)
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
		Message    string
		StatusCode int
	}{
		Message:    http.StatusText(statusCode),
		StatusCode: statusCode,
	}

	w.WriteHeader(statusCode)
	err = tmpl.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
