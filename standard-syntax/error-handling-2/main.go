package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

var ErrCustom = errors.New("custom error")

func main() {
	e1 := echo.NewHTTPError(http.StatusUnauthorized, "no user")
	e2 := errors.Wrap(e1, "wrap error")

	switch v := errors.Cause(e2).(type) {
	case *echo.HTTPError:
		fmt.Println("echo http error ", v.Code)
	default:
		fmt.Println("normal error")
	}

	e3 := fmt.Errorf("some error")
	e4 := errors.Wrap(e3, "wrap error")

	switch v := errors.Cause(e4).(type) {
	case *echo.HTTPError:
		fmt.Println("echo http error ", v.Code)
	default:
		fmt.Println("normal error", v.Error())
	}
}
