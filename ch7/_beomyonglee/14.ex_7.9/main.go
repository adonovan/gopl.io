package main

import (
	column "gopl.io/ch7/_beomyonglee/13.ex_7.8"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var people = []column.Person{
	{"Alice", 20},
	{"Bob", 12},
	{"Bob", 20},
	{"Alice", 12},
}

var html = template.Must(template.New("people").Parse(`
<html>
<body>
<table>
	<tr>
		<th><a href="?sort=name">name</a></th>
		<th><a href="?sort=age">age</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Name}}</td>
		<td>{{.Age}}</td>
	</td>
{{end}}
</body>
</html>
`))

func main() {
	c := column.NewByColumns(people, 2)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "age":
			c.Select(c.LessAge)
		case "name":
			c.Select(c.LessName)
		}
		sort.Sort(c)
		err := html.Execute(w, people)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
