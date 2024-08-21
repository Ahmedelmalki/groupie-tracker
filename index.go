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
	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations)

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("static/index.html")
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

// func handleProfile(w http.ResponseWriter, r *http.Request) {
// 	var artists []Artist
// 	var locations []LocationEntry
// 	var dates []DateEntry
// 	var relations []RelationEntry
// 	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
// 	fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
// 	fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
// 	fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations)
// 	tmpl, err := template.ParseFiles("static/profile.html")
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	err = tmpl.ExecuteTemplate(w, "index.html", struct {
// 		Artists   []Artist
// 		Locations []LocationEntry
// 		Dates     []DateEntry
// 		Relations []RelationEntry
// 	}{
// 		Artists:   artists,
// 		Locations: locations,
// 		Dates:     dates,
// 		Relations: relations,
// 	})
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

// Handler for the profile page
func handleProfile(w http.ResponseWriter, r *http.Request) {
	// Get the artist ID from the query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Artist ID not provided", http.StatusBadRequest)
		return
	}
	// Fetch the artist's data based on the ID
	var artist Artist
	err := fetchData("https://groupietrackers.herokuapp.com/api/artists/"+idStr, &artist)
	if err != nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Parse the profile HTML template file
	tmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the artist's data
	err = tmpl.ExecuteTemplate(w, "profile.html", artist)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
