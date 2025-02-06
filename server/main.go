package server

import (
	hands "RunShare/server/handlers"
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
	e.GET("/run", hands.RunHandler)
	e.POST("/run/:id", hands.GetRunById)

	e.Logger.Fatal(e.Start(":8080"))
}

func mainHandler(c echo.Context) error {
	return c.Render(200, "index", nil)
}
