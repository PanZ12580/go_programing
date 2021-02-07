/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

const Template = `{{.TotalCount}} issues
{{range .Items}}------------------------------------------
Numbers:	{{.Number}}
User:	{{.User.Login}} 
Title:	{{.Title | printf "%.62s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}
`

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items          []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
