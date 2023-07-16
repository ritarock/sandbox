package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name": "Dolly",
		})
	}).Name = "foobar"

	e.Logger.Fatal(e.Start(":8080"))
}
