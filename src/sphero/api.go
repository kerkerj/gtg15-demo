package main

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

  "github.com/thoas/stats"
)

func NewApi(sphero Sphero) Api {
	return &ApiEcho{sphero}
}

type ApiEcho struct {
	s Sphero
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

  // Serve static files
	e.Static("/", "./public")

  // stats https://github.com/thoas/stats
	s := stats.New()
	e.Use(s.Handler)
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})

	// ping
	e.Get("/ping", pingHandler())
	// Set RGB
	e.Get("/rgb/:rgb", rgbHandler(api))
	// SetStabilization
	e.Get("/stable/:bool", stableHandler(api))
	// SetBackLED
	e.Get("/backled/:level", backLedHandler(api))
	// SetHeading
	e.Get("/heading/:heading", headingHandler(api))
	// Roll
	e.Get("/roll/:speed/head/:heading", rollHandler(api))
  // SetRotationRate
  e.Get("/rotate_rate/:level", rotationRateHandler(api))

	return e
}

func pingHandler() echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.String(http.StatusOK, "pong")
	}
}

func rgbHandler(api *ApiEcho) echo.HandlerFunc {
	return func(c *echo.Context) error {
		rgb := c.Param("rgb")

		rgbRegexp := regexp.MustCompile(`^(\d+),(\d+),(\d+)$`)

		if !rgbRegexp.MatchString(rgb) {
			return c.String(http.StatusBadRequest, "error")
		}

		match := rgbRegexp.FindStringSubmatch(rgb)
		rr, _ := strconv.Atoi(match[1])
		gg, _ := strconv.Atoi(match[2])
		bb, _ := strconv.Atoi(match[3])
		api.s.SetRGB(uint8(rr), uint8(gg), uint8(bb))

		return c.String(http.StatusOK, "success")
	}
}

func stableHandler(api *ApiEcho) echo.HandlerFunc {
	return func(c *echo.Context) error {
		setStable, err := strconv.ParseBool(c.Param("bool"))

		if err != nil {
			return c.String(http.StatusBadRequest, "error")
		}

		api.s.SetStabilization(setStable)
		return c.String(http.StatusOK, "success")
	}
}

func backLedHandler(api *ApiEcho) echo.HandlerFunc {
	return func(c *echo.Context) error {
		level, err := strconv.Atoi(c.Param("level"))

		if err != nil {
			return c.String(http.StatusBadRequest, "error")
		}

		api.s.SetBackLED(uint8(level))
		return c.String(http.StatusOK, "success")
	}
}

func headingHandler(api *ApiEcho) echo.HandlerFunc {
	return func(c *echo.Context) error {
		heading, err := strconv.Atoi(c.Param("heading"))

		if err != nil {
			return c.String(http.StatusBadRequest, "error")
		}

		api.s.SetHeading(uint16(heading))
		return c.String(http.StatusOK, "success")
	}
}

func rollHandler(api *ApiEcho) echo.HandlerFunc {
	return func(c *echo.Context) error {
    speed, _ := strconv.Atoi(c.Param("speed"))
    heading, _ := strconv.Atoi(c.Param("heading"))

    api.s.Roll(uint8(speed), uint16(heading))
    return c.String(http.StatusOK, "success")
	}
}

func rotationRateHandler(api *ApiEcho) echo.HandlerFunc {
	return func(c *echo.Context) error {
		level, err := strconv.Atoi(c.Param("level"))

		if err != nil {
			return c.String(http.StatusBadRequest, "error")
		}

		api.s.SetRotationRate(uint8(level))
		return c.String(http.StatusOK, "success")
	}
}
