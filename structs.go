package main

// Struct to represent an individual artist
type Artist struct {
	Id      int      `json:"id"`
	Image   string   `json:"image"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

// struct to represent an individual location
type LocationEntry struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Locations struct {
	Index []LocationEntry `json:"index"`
}

// Struct to represent each date entry
type DateEntry struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Struct to represent the whole response if needed
type DatesResponse struct {
	Index []DateEntry `json:"index"`
}

// Struct to represent the relation entry
type RelationEntry struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Struct to represent the whole response if needed
type RelationResponse struct {
	Index []RelationEntry `json:"index"`
}
