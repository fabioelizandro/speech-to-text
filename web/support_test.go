package web_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func makeRequest(handler http.Handler, r *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, r)

	if response.Code == http.StatusSeeOther {
		return makeRequest(handler, httptest.NewRequest("GET", response.Header().Get("Location"), strings.NewReader("")))
	}

	return response
}

func createServer(handler http.Handler) func(r *http.Request) *httptest.ResponseRecorder {
	return func(r *http.Request) *httptest.ResponseRecorder {
		return makeRequest(handler, r)
	}
}

func submitForm(form *goquery.Selection, inputValues map[string]string) *http.Request {
	// hidden inputs and other than text input types are not handled... yet

	data := url.Values{}
	form.Find("input").Each(func(_ int, input *goquery.Selection) {
		if name, exists := input.Attr("name"); exists {
			data.Set(name, inputValues[name])
		}
	})

	for key := range inputValues {
		if !data.Has(key) {
			panic(fmt.Errorf("input value `%s` does not exist in the HTML form #%s", key, form.AttrOr("id", "")))
		}
	}

	request := httptest.NewRequest(
		strings.ToUpper(form.AttrOr("method", "INVALID")),
		path.Join("/", form.AttrOr("action", "invalid")),
		strings.NewReader(data.Encode()),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return request
}
