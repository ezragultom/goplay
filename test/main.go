package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func adminIndex(c echo.Context) error {

	delete_z := c.Param("delete_z")
	fmt.Println("delete_z", delete_z)

	get_z := c.Param("get_z")
	fmt.Println("get_z", get_z)

	post_z := c.Param("post_z")
	fmt.Println("post_z", post_z)

	return c.String(http.StatusOK, "Hello, Admin!")
}

func stu(c echo.Context) error {
	delete_z := c.Param("delete_z")
	fmt.Println("delete_z", delete_z)

	get_z := c.Param("get_z")
	fmt.Println("get_z", get_z)

	post_z := c.Param("post_z")
	fmt.Println("post_z", post_z)

	return c.String(http.StatusOK, "Hello, Admin!")
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	admin := e.Group("/admin")

	admin.GET("/test/:get_z", adminIndex)
	admin.POST("/test/:post_z", stu)
	admin.DELETE("/test/:delete_z", adminIndex)

	admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "123" {
			return true, nil
		}
		return false, nil
	}))

	e.Logger.Fatal(e.Start(":1323"))
}
