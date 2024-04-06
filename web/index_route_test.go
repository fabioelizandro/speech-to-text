package web_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/fabioelizandro/speech-to-text/store"
	"github.com/fabioelizandro/speech-to-text/web"
	"github.com/stretchr/testify/assert"
)

func Test_index_route(t *testing.T) {
	t.Run("renders index page with correct title", func(t *testing.T) {
		serve := web.New("../templates", store.NewInMemoryAudioStore())

		response := httptest.NewRecorder()
		serve.ServeHTTP(response, httptest.NewRequest("GET", "/", strings.NewReader("")))

		doc, err := goquery.NewDocumentFromReader(response.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Speech To Text", doc.Find("h1").Text())
	})
}
