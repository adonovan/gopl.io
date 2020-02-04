package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var listHTML = template.Must(template.New("list").Parse(`
<html>
<body>
<table>
	<tr>
		<th>item</th>
		<th>price</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td>{{$k}}</td>
		<td>{{$v}}</td>
	</tr>
{{end}}
</table>
</body>
</html>
`))

type PriceDB struct {
	sync.Mutex
	db map[string]int
}

func (p *PriceDB) Create(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "No integer price given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; ok {
		http.Error(w, fmt.Sprintf("%s already exists", item), http.StatusBadRequest)
		return
	}

	p.Lock()
	if p.db == nil {
		p.db = make(map[string]int, 0)
	}
	p.db[item] = price
	p.Unlock()
}

func (p *PriceDB) Update(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "No integer price given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	p.db[item] = price
	p.Unlock()
}

func (p *PriceDB) Delete(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}
	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	delete(p.db, item)
	p.Unlock()
}

func (p *PriceDB) Read(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	fmt.Fprintf(w, "%s: %d\n", item, p.db[item])
	p.Unlock()
}

func (p *PriceDB) List(w http.ResponseWriter, r *http.Request) {
	p.Lock()
	listHTML.Execute(w, p.db)
	p.Unlock()
}

func main() {
	db := &PriceDB{}
	db.db = make(map[string]int, 0)
	db.db["shoe"] = 100
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	http.HandleFunc("/list", db.List)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
$ go build gopl.io/ch1/_BeomyongLee/fetch
$ ./fetch http://localhost:8080/create\?item\=shoes\&price\=1
$ ./fetch http://localhost:8080/read\?item\=shoes
shoes: 1
$ ./fetch http://localhost:8080/update\?item\=shoes\&price\=10
$ ./fetch http://localhost:8080/read\?item\=shoes
shoes: 10
$  ./fetch http://localhost:8080/delete\?item\=shoes
$ ./fetch http://localhost:8080/read\?item\=shoes
shoes doesn't exist
*/
