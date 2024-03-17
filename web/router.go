package web

import (
	"fmt"
	"net/http"

	"github.com/fabioelizandro/speech-to-text/webtmpl"
	"github.com/julienschmidt/httprouter"
)

func Router(renderer webtmpl.Renderer) http.Handler {
	router := httprouter.New()
	router.GET("/", handleRouteError(indexRouter(renderer)))

	return router
}

type routeWithError func(http.ResponseWriter, *http.Request, httprouter.Params) error

// small interface adaptor so routes can delegate error handling to caller
func handleRouteError(route routeWithError) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		err := route(writer, request, params)
		if err != nil {
			fmt.Printf("Something went wrong with request: %+v\n", err)
			writer.WriteHeader(500)
		}
	}
}
