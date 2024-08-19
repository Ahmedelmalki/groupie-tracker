package main

import (
	"net/http"
	"text/template"
)

type ArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Fetch artists data
	var artists []Artist
	var locations []LocationEntry
	var dates []DateEntry
	var relations []RelationEntry
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations)
	fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the fetched data
	err = tmpl.ExecuteTemplate(w, "index.html", struct {
		Artists   []Artist
		Locations []LocationEntry
		Dates     []DateEntry
		Relations []RelationEntry
	}{
		Artists:   artists,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
