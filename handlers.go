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
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var artists []Artist
	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errors(w, http.StatusMethodNotAllowed)
		return
	}
	// Get the artist ID from the query parameters
	idStr := r.URL.Query().Get("id")
	patren := `\d+$`
	re := regexp.MustCompile(patren)
	if !re.MatchString(idStr) {
		errors(w, http.StatusNotFound)
		return
	}
	if idStr == "" {
		errors(w, http.StatusMethodNotAllowed)
		return
	}
	var artist Artist
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &artist)
	if err != nil {
		fmt.Println(1)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/locations/"+idStr, &artist.Locations)
	if err != nil {
		fmt.Println(2)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/dates/"+idStr, &artist.Dates)
	if err != nil {
		fmt.Println(3)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = fetchData("https://groupietrackers.herokuapp.com/api/relation/"+idStr, &artist.Relations)
	if err != nil {
		fmt.Println(4)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if artist.Name == "" {
		errors(w, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "profile.html", artist)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
}

func errors(w http.ResponseWriter, statusCode int) {
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
	fmt.Println(data.Message)
}
