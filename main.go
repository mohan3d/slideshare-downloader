package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/mohan3d/slideshare-go/slideshare"
)

func fileName(url string) string {
	return filepath.Base(url) + ".pdf"
}

func quality(q string) slideshare.Quality {
	// high is the default quality
	v := slideshare.QualityFull

	switch q {
	case "normal":
		v = slideshare.QualityNormal
	case "low":
		v = slideshare.QualitySmall
	}
	return v
}

func fetchHandler(c *gin.Context) {
	query := c.Request.URL.Query()
	u := query.Get("url")
	q := query.Get("quality")

	var buf bytes.Buffer
	err := slideshare.DefaultSlideshareDownloader.Download(
		u,
		quality(q),
		&buf,
	)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.tmpl.html", nil)
	} else {
		c.Header("Content-Disposition", "attachment; filename="+fileName(u))
		c.Data(http.StatusOK, "application/pdf", buf.Bytes())
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/fetch", fetchHandler)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
