package main

import (
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

type Album struct {
	ID     string `json:"DUDUD" jsonapi:"primary,cart"`
	Name   string `jsonapi:"attr,name"`
	Artist Artist `jsonapi:"attr,artist"`
}
type Artist struct {
	Name string `jsonapi:name`
	Age  int    `jsonapi:age`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		jsonapi.MarshalPayload(c.Response(), albumList())
		return c.JSON(http.StatusOK, c.Response())

	})
	e.Logger.Fatal(e.Start(":1323"))
}

func albumList() []*Album {
	a1 := Album{"123", "allbum1", Artist{Name: "stu", Age: 17}}
	a2 := Album{"456", "allbum2", Artist{Name: "ezra", Age: 18}}
	albums := []*Album{&a1, &a2}
	return albums
}
