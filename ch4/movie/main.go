// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 108.

// Movie prints Movies as JSON.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/bxcodec/faker/v3"
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
	// ...
}

type Person struct{
	FirstName string 
	TitleMale string 
	CCNumber int `json:"ccnumbers"`
	Sibling []string `json:"siblings"` 
}

var CCNumber, _ = strconv.ParseInt(faker.CCNumber(), 0, 64)


var personInfo = []Person{
	{	FirstName: faker.FirstName(),
		TitleMale: faker.TitleMale(),
		CCNumber: int(CCNumber),
		Sibling: []string{faker.Name(), faker.Name(),},
	},
}
//!-

func main() {
	
		//!+Marshal
		// _, err := json.Marshal(movies)
		// if err != nil {
		// 	log.Fatalf("JSON marshaling failed: %s", err)
		// }
		//fmt.Printf("personInfo: %s \n", personInfo)
		//!-Marshal
		personData, err := json.Marshal(personInfo)
		if err != nil {
			log.Fatalf("JSON marshling failed: %s", err)
		}
		fmt.Printf("personData: %s \n",personData)

		data, err := json.MarshalIndent(personInfo, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("Pretty JSON: %s\n", data)

	
		//!+MarshalIndent
		// data, err := json.MarshalIndent(movies, "", "    ")
		// if err != nil {
		// 	log.Fatalf("JSON marshaling failed: %s", err)
		// }
		// fmt.Printf("%s\n", data)
		// //!-MarshalIndent

		// //!+Unmarshal
		// var titles []struct{ Title string }
		// if err := json.Unmarshal(data, &titles); err != nil {
		// 	log.Fatalf("JSON unmarshaling failed: %s", err)
		// }
		// fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
		// //!-Unmarshal
	
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
