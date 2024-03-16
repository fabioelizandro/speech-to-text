package web

import (
	"net/http"

	"github.com/fabioelizandro/speech-to-text/webtmpl"
	"github.com/julienschmidt/httprouter"
)

func helloRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	w.Header().Add("Content-Type", "text/html")

	return webtmpl.RenderHTML(w, webtmpl.HelloTemplate{
		Name: ps.ByName("name"),
	})
}
