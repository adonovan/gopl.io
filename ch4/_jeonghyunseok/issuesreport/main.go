// Issuereport 는 이슈에 대해 리포트를 해준다.
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"gopl.io/ch4/github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nnoMust version")
	noMust()
}

// https://blog.gopheracademy.com/advent-2017/using-go-templates/
func noMust() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

/*
//!+output
go build -o issuesreport.exe
issuesreport.exe repo:golang/go is:open json decoder

13 issues:
----------------------------------------
Number: 5680
User:   eaigner
Title:  encoding/json: set key converter on en/decoder
Age:    750 days
----------------------------------------
Number: 6050
User:   gopherbot
Title:  encoding/json: provide tokenizer
Age:    695 days
----------------------------------------
...
//!-output
*/

// 44 issues:
// -----------------------------
// Number: 36225
// User: dsnet
// Title: encoding/json: the Decoder.Decode API lends itself to misuse
// Age: 4 days
// -----------------------------
// Number: 33416
// User: bserdar
// Title: encoding/json: This CL adds Decoder.InternKeys
// Age: 144 days
// -----------------------------
// Number: 34647
// User: babolivier
// Title: encoding/json: fix byte counter increments when using decoder.To
// Age: 83 days
// -----------------------------
// Number: 5901
// User: rsc
// Title: encoding/json: allow override type marshaling
// Age: 2350 days
// -----------------------------
// Number: 29035
// User: jaswdr
// Title: proposal: encoding/json: add error var to compare  the returned
// Age: 388 days
// -----------------------------
// Number: 34543
// User: maxatome
// Title: encoding/json: Unmarshal & json.(*Decoder).Token report differen
// Age: 89 days
// -----------------------------
// Number: 32779
// User: rsc
// Title: proposal: encoding/json: memoize strings during decode?
// Age: 181 days
// -----------------------------
// Number: 28923
// User: mvdan
// Title: encoding/json: speed up the decoding scanner
// Age: 396 days
// -----------------------------
// Number: 11046
// User: kurin
// Title: encoding/json: Decoder internally buffers full input
// Age: 1664 days
// -----------------------------
// Number: 34564
// User: mdempsky
// Title: go/internal/gcimporter: single source of truth for decoder logic
// Age: 87 days
// -----------------------------
// Number: 12001
// User: lukescott
// Title: encoding/json: Marshaler/Unmarshaler not stream friendly
// Age: 1603 days
// -----------------------------
// Number: 31789
// User: mgritter
// Title: encoding/json: provide a way to limit recursion depth
// Age: 236 days
// -----------------------------
// Number: 30301
// User: zelch
// Title: encoding/xml: option to treat unknown fields as an error
// Age: 308 days
// -----------------------------
// Number: 33854
// User: Qhesz
// Title: encoding/json: unmarshal option to treat omitted fields as null
// Age: 118 days
// -----------------------------
// Number: 31701
// User: lr1980
// Title: encoding/json: second decode after error impossible
// Age: 241 days
// -----------------------------
// Number: 26946
// User: deuill
// Title: encoding/json: clarify what happens when unmarshaling into a non
// Age: 498 days
// -----------------------------
// Number: 16212
// User: josharian
// Title: encoding/json: do all reflect work before decoding
// Age: 1272 days
// -----------------------------
// Number: 33835
// User: Qhesz
// Title: encoding/json: unmarshalling null into non-nullable golang types
// Age: 119 days
// -----------------------------
// Number: 6647
// User: btracey
// Title: x/tools/cmd/godoc: display type kind of each named type
// Age: 2252 days
// -----------------------------
// Number: 27179
// User: lavalamp
// Title: encoding/json: no way to preserve the order of map keys
// Age: 487 days
// -----------------------------
// Number: 28143
// User: arp242
// Title: proposal: encoding/json: add "readonly" tag
// Age: 438 days
// -----------------------------
// Number: 22752
// User: buyology
// Title: proposal: encoding/json: add access to the underlying data causi
// Age: 768 days
// -----------------------------
// Number: 28189
// User: adnsv
// Title: encoding/json: confusing errors when unmarshaling custom types
// Age: 436 days
// -----------------------------
// Number: 22480
// User: blixt
// Title: proposal: encoding/json: add omitnil option
// Age: 786 days
// -----------------------------
// Number: 33714
// User: flimzy
// Title: proposal: encoding/json: Opt-in for true streaming support
// Age: 126 days
// -----------------------------
// Number: 14750
// User: cyberphone
// Title: encoding/json: parser ignores the case of member names
// Age: 1383 days
// -----------------------------
// Number: 7872
// User: extemporalgenome
// Title: encoding/json: Encoder internally buffers full output
// Age: 2067 days
// -----------------------------
// Number: 30701
// User: LouAdrien
// Title: encoding/json: ignore tag "-" not working on embedded sub struct
// Age: 289 days
// -----------------------------
// Number: 20528
// User: jvshahid
// Title: net/http: connection reuse does not work happily with normal use
// Age: 937 days
// -----------------------------
// Number: 20754
// User: rsc
// Title: encoding/xml: unmarshal only processes first XML element
// Age: 914 days

// noMust version
// 44 issues:
// -----------------------------
// Number: 36225
// User: dsnet
// Title: encoding/json: the Decoder.Decode API lends itself to misuse
// Age: 4 days
// -----------------------------
// Number: 33416
// User: bserdar
// Title: encoding/json: This CL adds Decoder.InternKeys
// Age: 144 days
// -----------------------------
// Number: 34647
// User: babolivier
// Title: encoding/json: fix byte counter increments when using decoder.To
// Age: 83 days
// -----------------------------
// Number: 5901
// User: rsc
// Title: encoding/json: allow override type marshaling
// Age: 2350 days
// -----------------------------
// Number: 29035
// User: jaswdr
// Title: proposal: encoding/json: add error var to compare  the returned
// Age: 388 days
// -----------------------------
// Number: 34543
// User: maxatome
// Title: encoding/json: Unmarshal & json.(*Decoder).Token report differen
// Age: 89 days
// -----------------------------
// Number: 32779
// User: rsc
// Title: proposal: encoding/json: memoize strings during decode?
// Age: 181 days
// -----------------------------
// Number: 28923
// User: mvdan
// Title: encoding/json: speed up the decoding scanner
// Age: 396 days
// -----------------------------
// Number: 11046
// User: kurin
// Title: encoding/json: Decoder internally buffers full input
// Age: 1664 days
// -----------------------------
// Number: 34564
// User: mdempsky
// Title: go/internal/gcimporter: single source of truth for decoder logic
// Age: 87 days
// -----------------------------
// Number: 12001
// User: lukescott
// Title: encoding/json: Marshaler/Unmarshaler not stream friendly
// Age: 1603 days
// -----------------------------
// Number: 31789
// User: mgritter
// Title: encoding/json: provide a way to limit recursion depth
// Age: 236 days
// -----------------------------
// Number: 30301
// User: zelch
// Title: encoding/xml: option to treat unknown fields as an error
// Age: 308 days
// -----------------------------
// Number: 33854
// User: Qhesz
// Title: encoding/json: unmarshal option to treat omitted fields as null
// Age: 118 days
// -----------------------------
// Number: 31701
// User: lr1980
// Title: encoding/json: second decode after error impossible
// Age: 241 days
// -----------------------------
// Number: 26946
// User: deuill
// Title: encoding/json: clarify what happens when unmarshaling into a non
// Age: 498 days
// -----------------------------
// Number: 16212
// User: josharian
// Title: encoding/json: do all reflect work before decoding
// Age: 1272 days
// -----------------------------
// Number: 33835
// User: Qhesz
// Title: encoding/json: unmarshalling null into non-nullable golang types
// Age: 119 days
// -----------------------------
// Number: 6647
// User: btracey
// Title: x/tools/cmd/godoc: display type kind of each named type
// Age: 2252 days
// -----------------------------
// Number: 27179
// User: lavalamp
// Title: encoding/json: no way to preserve the order of map keys
// Age: 487 days
// -----------------------------
// Number: 28143
// User: arp242
// Title: proposal: encoding/json: add "readonly" tag
// Age: 438 days
// -----------------------------
// Number: 22752
// User: buyology
// Title: proposal: encoding/json: add access to the underlying data causi
// Age: 768 days
// -----------------------------
// Number: 28189
// User: adnsv
// Title: encoding/json: confusing errors when unmarshaling custom types
// Age: 436 days
// -----------------------------
// Number: 22480
// User: blixt
// Title: proposal: encoding/json: add omitnil option
// Age: 786 days
// -----------------------------
// Number: 33714
// User: flimzy
// Title: proposal: encoding/json: Opt-in for true streaming support
// Age: 126 days
// -----------------------------
// Number: 14750
// User: cyberphone
// Title: encoding/json: parser ignores the case of member names
// Age: 1383 days
// -----------------------------
// Number: 7872
// User: extemporalgenome
// Title: encoding/json: Encoder internally buffers full output
// Age: 2067 days
// -----------------------------
// Number: 30701
// User: LouAdrien
// Title: encoding/json: ignore tag "-" not working on embedded sub struct
// Age: 289 days
// -----------------------------
// Number: 20528
// User: jvshahid
// Title: net/http: connection reuse does not work happily with normal use
// Age: 937 days
// -----------------------------
// Number: 20754
// User: rsc
// Title: encoding/xml: unmarshal only processes first XML element
// Age: 914 days
