package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func bar() error {
	if err := foo(); err != nil {
		return errors.WithMessage(foo(), "bar failed")
	}
	return nil
}

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func main() {
	err := bar()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("data not found,  %+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}
}
