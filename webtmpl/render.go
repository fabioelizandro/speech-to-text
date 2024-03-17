package webtmpl

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed views/*.gohtml
var htmlViews embed.FS

type Renderer interface {
	RenderHTML(w io.Writer, templ Template) error
}

type EmbeddedRenderer struct {
}

func NewEmbeddedRenderer() *EmbeddedRenderer {
	return &EmbeddedRenderer{}
}

func (r *EmbeddedRenderer) RenderHTML(w io.Writer, templ Template) error {
	templates := template.Must(template.ParseFS(htmlViews, "views/_*.gohtml", fmt.Sprintf("views/%s", templ._name())))

	return templates.ExecuteTemplate(w, templ._name(), templ)
}

type FileRenderer struct {
}

func NewFileRenderer() *FileRenderer {
	return &FileRenderer{}
}

func (r *FileRenderer) RenderHTML(w io.Writer, templ Template) error {
	templates := template.Must(template.ParseGlob("webtmpl/views/_*.gohtml"))
	templates = template.Must(templates.ParseGlob(fmt.Sprintf("webtmpl/views/%s", templ._name())))

	return templates.ExecuteTemplate(w, templ._name(), templ)
}
