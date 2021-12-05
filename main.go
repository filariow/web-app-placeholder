package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/filariow/web-app-placeholder/assets"
	"github.com/filariow/web-app-placeholder/pkg/host"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	if err := setupHTTPServer(); err != nil {
		return err
	}
	return nil
}

func setupHTTPServer() error {
	r := gin.Default()

	t, err := template.ParseFS(assets.Templates, "templates/home-page.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(t)

	a := getAddress()
	rg := r.Group(a)
	rg.GET("/", getHomePage)
	rg.GET("/favicon.svg", func(c *gin.Context) {
		c.FileFromFS("public/favicon.svg", http.FS(assets.Public))
	})

	r.Run(getPort())

	return nil
}

func getHomePage(c *gin.Context) {
	i := host.Info()

	c.HTML(http.StatusOK, "home-page.tmpl", gin.H{
		"title": "I'm a Placeholder Web App",
		"host":  i,
	})
}

func getAddress() string {
	if a, ok := os.LookupEnv("URL_PATH"); ok {
		return a
	}

	return ""
}

func getPort() string {
	if val, ok := os.LookupEnv("PORT"); ok {
		return ":" + val
	}

	return ":8080"
}
