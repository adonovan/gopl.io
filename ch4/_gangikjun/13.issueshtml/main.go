package main

import (
	"html/template"
	"log"
	"os"

	"gopl.io/ch4/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
    <tr style="text-align: left;">
        <th>#</th>
        <th>State</th>
        <th>User</th>
        <th>Title</th>
    </tr>
    {{range .Items}}
    <tr>
        <td><a href="{{HTMLURL}}">{{.Number}}</a></td>
        <td>{{.State}}</td>
        <td><a href="{{.User>.HTMLURL}}">{{.User.Login}}</a></td>
        <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
</table>
`))

func main() {
	q := []string{
		"repo:golang/go",
		"is:open",
		"json",
		"decoder",
	}

	result, err := github.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
