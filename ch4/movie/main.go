// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 108.

// Movie prints Movies as JSON.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//!+
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	{Title: "Star Wars", Year: 1978, Color: true,
		Actors: []string{"Carrie Fisher", "Alec Ginnus", "Mark Hammel", "James Earl Jones (Voice)"}},
	// ...
}

//!-

func main() {
	{
		//!+Marshal
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n\n", data)
		//!-Marshal
	}

	{
		//!+MarshalIndent
		data, err := json.MarshalIndent(movies, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n\n", data)
		//!-MarshalIndent

		//!+Unmarshal
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}

		//Add by Doug to unmarshal the actors from the json
		//gets the actors and stores them in a slice of structs
		//!+Unmarshal
		var actors []struct{ Actors []string }
		if err := json.Unmarshal(data, &actors); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		//!-Unmarshal

		fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
		fmt.Println(actors) // Added by doug to print the actors

	}
}

/*
//!+output
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingr
id Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Ac
tors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"
Actors":["Steve McQueen","Jacqueline Bisset"]}]
//!-output
*/

/*
//!+indented
[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
    {
        "Title": "Cool Hand Luke",
        "released": 1967,
        "color": true,
        "Actors": [
            "Paul Newman"
        ]
    },
    {
        "Title": "Bullitt",
        "released": 1968,
        "color": true,
        "Actors": [
            "Steve McQueen",
            "Jacqueline Bisset"
        ]
    }
]
//!-indented
*/
