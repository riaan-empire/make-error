package main

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type Base struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type Data struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Base
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Pre(middleware.AddTrailingSlash())

	// open := e.Group("/api")

	e.GET("/", func(c echo.Context) error {
		resp := resp()

		return c.JSON(resp.StatusCode, resp)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func resp() Data {
	i := rand.Intn(13)
	resp := Data{}

	switch i {
	case 0:
		resp.Name = "RandomName"
		resp.Code = "supersecretawesomecode"
		resp.StatusCode = http.StatusOK
		resp.Message = "Success"
	case 1:
		resp.StatusCode = http.StatusBadGateway
		resp.Message = "Bad Gateway"
	case 2:
		resp.StatusCode = http.StatusNotFound
		resp.Message = "Not Found"
	case 3:
		resp.StatusCode = http.StatusUnauthorized
		resp.Message = "Unauthorized"
	case 4:
		resp.StatusCode = http.StatusPermanentRedirect
		resp.Message = "Redirected"
	case 5:
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad Request"
	case 6:
		resp.StatusCode = http.StatusGatewayTimeout
		resp.Message = "Gateway Timeout"
	default:
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal Server Error"
	}

	return resp
}
