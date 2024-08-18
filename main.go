package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Struct to hold the initial API response URLs
type APIResponse struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

// Struct to represent an individual artist
type Artist struct {
	Id      int      `json:"id"`
	Image   string   `json:"image"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

// Struct to hold the list of artists
type Artists []Artist

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject APIResponse
	// Unmarshal the response into the APIResponse struct
	json.Unmarshal(responseData, &responseObject)

	// Now, make an HTTP request to get the artists' data
	artistsResponse, err := http.Get(responseObject.Artists)
	if err != nil {
		log.Fatal(err)
	}
	artistsData, err := io.ReadAll(artistsResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var artists Artists
	err = json.Unmarshal(artistsData, &artists)
	if err != nil {
		log.Fatal(err)
	}
	// Print out the artist data as an example
	for _, artist := range artists {
		fmt.Printf("Artist ID: %d, Name: %s, Image: %s, Members: %v\n", artist.Id, artist.Name, artist.Image, artist.Members)
	}

	fmt.Println(responseObject.Artists)
	fmt.Println(responseObject.Locations)
	fmt.Println(responseObject.Dates)
	fmt.Println(responseObject.Relation)
}
