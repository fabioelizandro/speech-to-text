package webtmpl

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed views/*.gohtml
var htmlViews embed.FS

func RenderHTML(w io.Writer, templ Template) error {
	// We can use memoization for the template parsing logic
	templates := template.Must(template.ParseFS(htmlViews, "views/_*.gohtml", fmt.Sprintf("views/%s", templ._name())))

	return templates.ExecuteTemplate(w, templ._name(), templ)

}
