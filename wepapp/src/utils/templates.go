package utils

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplate() {
	templates = template.Must(template.ParseGlob("./views/*.html"))
}

func RenderTemplate(w http.ResponseWriter, template string, data any) {
	view := fmt.Sprintf("%s.html", template)
	if err := templates.ExecuteTemplate(w, view, data); err != nil {
		log.Fatal(err)
	}
}
