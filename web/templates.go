package web

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed views/*.gohtml
var htmlViews embed.FS

func renderHTML(w http.ResponseWriter, templateName string, data any) error {
	// We can use memoization for the template parsing logic
	templates := template.Must(template.ParseFS(htmlViews, "views/_*.gohtml", fmt.Sprintf("views/%s", templateName)))

	w.Header().Add("Content-Type", "text/html")
	return templates.ExecuteTemplate(w, templateName, data)
}
