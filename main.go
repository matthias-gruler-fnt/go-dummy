package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ( // injected upon building by goreleaser
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Println("==========================================================")
	fmt.Println("                       go-dummy")
	fmt.Println("----------------------------------------------------------")
	fmt.Printf(" Version    : %s\n", version)
	fmt.Printf(" Build Date : %s\n", date)
	fmt.Printf(" Git Commit : %s\n", commit)
	fmt.Println("==========================================================")

	ech := echo.New()

	ech.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	ech.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			struct {
				Status string
				Up     bool
			}{
				Status: "OK",
				Up:     true,
			},
		)
	})
	ech.GET("/items", func(c echo.Context) error {
		type item struct {
			ID    string
			Name  string
			Count int
		}
		type list struct {
			Title string
			Items []item
		}

		return c.JSON(
			http.StatusOK,
			list{
				Title: "Item list",
				Items: []item{
					{
						ID:    "i1",
						Name:  "Item 1",
						Count: 42,
					},
					{
						ID:    "i2",
						Name:  "Item Numero Due",
						Count: 123,
					},
					{
						ID:    "i3",
						Name:  "Item #3",
						Count: 5679645,
					},
					{
						ID:    "i4",
						Name:  "All your base are belong to us!",
						Count: 0,
					},
				},
			},
		)
	})
	ech.GET("/stderror", func(c echo.Context) error {
		return errors.New("error")
	})
	ech.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusTooManyRequests, "bad")
	})
	ech.Logger.Fatal(ech.Start(":8000"))
}
