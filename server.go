package main

import (
	"github.com/labstack/echo"
	"net/http"
	"myecho/controller"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	//e.POST("/users",saveUser)
	//e.GET("/users/:id",getUser)
	//e.PUT("/users/:id",updateUser)
	//e.DELETE("/users/:id",deleteUser)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users",controller.SaveUser)
	e.POST("/users/:id",controller.UpdateUser)
	e.DELETE("/users/:id",controller.DeleteUser)
	e.Logger.Fatal(e.Start(":1223"))
}
