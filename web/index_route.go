package web

import (
	"net/http"

	"github.com/fabioelizandro/speech-to-text/webtmpl"
	"github.com/julienschmidt/httprouter"
)

func indexRouter(renderer webtmpl.Renderer) routeWithError {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
		w.Header().Add("Content-Type", "text/html")

		return renderer.RenderHTML(w, webtmpl.IndexTemplate{})
	}
}
