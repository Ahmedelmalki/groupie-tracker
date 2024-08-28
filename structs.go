package main

// Struct to represent an individual artist
type Artist struct {
	Id           int       `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int       `json:"creationDate"`
	FristAlbum   string    `json:"firstAlbum"`
	Locations    Locations `json:"-"`
	Dates        Dates     `json:"-"`
	Relations    Relations `json:"-"`
}

type Locations struct {
	Locations []string `json:"locations"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

type Relations struct {
	Relation map[string][]string `json:"datesLocations"`
}
