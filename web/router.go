package web

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fabioelizandro/speech-to-text/assert"
	"github.com/fabioelizandro/speech-to-text/store"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func New(templatesDir string, store store.AudioStore) *gin.Engine {
	r := gin.Default()
	r.HTMLRender = renderer(templatesDir)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{})
	})

	r.POST("/audio-upload", func(c *gin.Context) {
		file := assert.Must(c.FormFile("audio-file"))

		assert.NoErr(store.SaveAudio(
			assert.Must(file.Open()),
			file.Filename,
		))

		c.Redirect(http.StatusFound, "/audios")
	})

	r.GET("/audios", func(c *gin.Context) {
		files := assert.Must(store.ListSavedFiles())

		c.HTML(200, "audios", gin.H{
			"Files": files,
		})
	})

	return r
}

func renderer(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts := assert.Must(filepath.Glob(filepath.Join(templatesDir, "/layouts/*.gohtml")))
	pages := assert.Must(filepath.Glob(filepath.Join(templatesDir, "/pages/*.gohtml")))

	for _, page := range pages {
		r.AddFromFiles(
			strings.TrimSuffix(
				filepath.Base(page),
				filepath.Ext(page),
			),
			append(layouts, page)...,
		)
	}

	return r
}
