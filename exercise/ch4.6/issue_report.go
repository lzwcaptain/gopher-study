package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
Number: {{.Number}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

type Issue struct {
	TotalCount int64     `json:"totalCount"`
	Items      []string  `json:"items"`
	Number     int64     `json:"number"`
	Title      string    `json:"title"`
	CreateAt   time.Time `json:"createAt"`
}

func main() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatalln(err)
	}
	result := Issue{
		TotalCount: 10,
		Items:      []string{"a", "b"},
		Number:     4,
		Title:      "github",
		CreateAt:   time.Now().Add(time.Duration(2 * time.Hour)),
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatalln(err)
	}
}
