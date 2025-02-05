package hands

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Run struct {
	RunId, CreatorId, Distance int
	Name, Details, Rating      string
}

func getRunJSON() []Run {
	var runs []Run

	return runs
}

func createRun(name, detail, rating string, distance int) {

}
func RunHandler(c echo.Context) error {
	return c.Render(200, "Run", nil)
}

func NewRuns(c echo.Context) error {
	r := new(Run)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	run := Run{
		Name:     r.Name,
		Details:  r.Details,
		Distance: r.Distance,
	}
	return c.JSON(http.StatusAccepted, r)
}
