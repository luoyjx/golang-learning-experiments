package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

func HandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if err = next(c); err != nil {
			c.Error(err)
		}
		url := c.Path()

		fmt.Println(url)

		return
	}
}

func main() {
	e := echo.New()
	e.GET("/:fullname", func(c echo.Context) error {
		name := c.Param("fullname")
		return c.String(http.StatusOK, fmt.Sprintf("Hello get, %s", name))
	})
	e.DELETE("/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
	})
	e.GET("/table-api/base/:guid", func(c echo.Context) error {
		name := c.Param("fullname")

		tableApiURL, _ := url.Parse("http://svc-table:9001")

		fmt.Println(c.Path())

		fmt.Println(c.Request().URL.Path)
		fmt.Println(strings.ReplaceAll(c.Request().URL.Path, "/table-api", ""))
		fmt.Println(fmt.Sprintf("%s%s", tableApiURL.String(), strings.ReplaceAll(c.Request().URL.Path, "/table-api", "")))

		return c.String(http.StatusOK, fmt.Sprintf("Hello get, %s", name))
	}, HandlerFunc)
	e.Logger.Fatal(e.Start(":1323"))
}
