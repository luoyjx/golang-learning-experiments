package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func A(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("A")

		return next(c)
	}
}

func B(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("B")

		return next(c)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}, A, B)
	e.Logger.Fatal(e.Start(":1323"))
}
