package main

import (
	"net/http"
	"strconv"
  "regexp"

  "github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func NewApi(sphero Sphero) Api {
	return &ApiEcho{sphero}
}

type ApiEcho struct {
	Gundam Sphero
}

type Response struct {
  Success     string  `json:"success"`
  StatusText  string  `json:"status_text"`
}

type Api interface {
	Handler() http.Handler
}

func (api *ApiEcho) Handler() http.Handler {
  // Echo instance
	e := echo.New()

  // Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

  // Set RGB
  e.Get("/rgb/:rgb", func(c *echo.Context) error {
    // By name
    rgb := c.Param("rgb")

    rgbRegexp := regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)

    if !rgbRegexp.MatchString(rgb) {
			return c.JSON(400, &Response { Success:  "false", StatusText: "Invalid format of rgb"})
		}

    match := rgbRegexp.FindStringSubmatch(rgb)
    rr, _ := strconv.Atoi(match[1])
  	gg, _ := strconv.Atoi(match[2])
  	bb, _ := strconv.Atoi(match[3])
  	api.Gundam.SetRGB((uint8)(rr), (uint8)(gg), (uint8)(bb))

    return c.JSON(http.StatusOK, &Response { Success:  "true", StatusText: "success" })
  })

  // TODO: SetSpin, SetLocate

	return e
}
