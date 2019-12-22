// Movie 는 영화들 정보를 JSON 으로 출력해준다.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:color,omitempty"` // 값이 없으면 아예 필드를 빼버린다
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Burgman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {

	{
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("fail: %v", err)
		}
		fmt.Printf("%s\n", data)
	}

	{
		// data, err 는 다른 스코프이다. 새로이 선언 필요
		data, err := json.MarshalIndent(movies, "↘", "----")
		if err != nil {
			log.Fatalf("fail: %v", err)
		}
		fmt.Printf("%s\n", data)

		var titles []struct{ Title string }
		if err := json.Unmarshal(data, titles); err != nil {
			log.Fatalf("fail: %v", err)
		}
		fmt.Println(titles)
		fmt.Println("--")
		fmt.Printf("%#v\n", titles)
	}
}

// go run main.go
