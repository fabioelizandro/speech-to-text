package web_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/fabioelizandro/speech-to-text/web"
	"github.com/stretchr/testify/assert"
)

func Test_hello_route(t *testing.T) {
	t.Run("renders hello with user name", func(t *testing.T) {
		serve := createServer(web.Router())

		response := serve(httptest.NewRequest("GET", "/hello/john", strings.NewReader("")))

		doc, err := goquery.NewDocumentFromReader(response.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Hello john =)", doc.Find("h1").Text())
	})
}
