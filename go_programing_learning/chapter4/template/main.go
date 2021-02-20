package main

import (
	"go_programing/chapter4/json/github"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/getIssues", getIssues)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	button := "<a href='/getIssues'>getIssues</a>"
	btn, err := template.New("issuesList").
		Parse(button)
	if err != nil {
		log.Fatal(err)
	}
	btn.Execute(w, nil)
}

func getIssues(w http.ResponseWriter, r *http.Request) {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": github.DaysAgo}).
		Parse(`<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})

	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(w, resp); err != nil {
		log.Fatal(err)
	}
}
