package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func downloadHandler(c echo.Context) error {
	url := c.FormValue("url")
	go Download(url)
	return c.String(http.StatusOK, "Downloading "+url)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.POST("/download", downloadHandler)
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
