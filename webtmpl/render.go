package webtmpl

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed views/*.gohtml
var htmlViews embed.FS

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderHTML(w io.Writer, templ Template) error {
	templates := template.Must(template.ParseFS(htmlViews, "views/_*.gohtml", fmt.Sprintf("views/%s", templ._name())))

	return templates.ExecuteTemplate(w, templ._name(), templ)
}
