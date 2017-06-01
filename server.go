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

	//path param
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users",controller.SaveUser)
	e.POST("/users/:id",controller.UpdateUser)
	e.DELETE("/users/:id",controller.DeleteUser)

	//query param
	e.GET("/show",controller.Show)

	//form param application/x-www-form-urlencoded
	e.POST("/save",controller.Save)

	//form multipart/form-data
	e.POST("/save",controller.SaveAvatar)



	e.Logger.Fatal(e.Start(":1323"))
}
