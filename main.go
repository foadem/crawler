package main

import (
	"html/template"
	"log"
	"net/http"

	"jurino.ir/project/controller"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	controller.CrawlTechnolife()
	httpserver()
}

func httpserver() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "HomePage")
}
