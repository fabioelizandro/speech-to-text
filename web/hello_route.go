package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type helloRouteData struct {
	Name string
}

func helloRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	return renderHTML(w, "hello.gohtml", helloRouteData{
		Name: ps.ByName("name"),
	})
}
