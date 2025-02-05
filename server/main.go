package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("app/views/*/*.html")),
	}
}

func Init() {
	e := echo.New()
	e.Static("/static", "app")
	e.Renderer = NewTemplate()
	e.GET("/", mainHandler)
	e.GET("/run", runHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func mainHandler(c echo.Context) error {
	return c.Render(200, "index", nil)
}

func runHandler(c echo.Context) error {
	return c.Render(200, "Run", nil)
}
