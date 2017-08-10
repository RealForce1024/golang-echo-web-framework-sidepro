package main

import (
	"github.com/labstack/echo"
	"net/http"
	"myecho/controller"
	"myecho/model"
	"github.com/labstack/echo/middleware"
)

var e = echo.New()

func init() {
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func main() {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	//e.POST("/users",saveUser)
	//e.GET("/users/:id",getUser)
	//e.PUT("/users/:id",updateUser)
	//e.DELETE("/users/:id",deleteUser)

	//path param
	e.GET("/users/:id", controller.GetUser)
	//e.POST("/users", controller.SaveUser)
	e.POST("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	//query param
	e.GET("/show", controller.Show)

	//form param application/x-www-form-urlencoded
	e.POST("/save", controller.Save)

	//form multipart/form-data
	e.POST("/save", controller.SaveAvatar)

	//$~  myecho [master] curl -F "name=fqc" -F "email=gomaster_me@sina.com" http://localhost:1323/users
	//{"Id":0,"name":"fqc","email":"gomaster_me@sina.com"}%
	e.POST("/users", func(c echo.Context) error {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
		//return c.XML(http.StatusCreated, u)
		//☁  Desktop  curl -F "name=Joe Smith" -F "email=xxx@sina.com" http://localhost:1323/users
		//<?xml version="1.0" encoding="UTF-8"?>
		//<User><Id>0</Id><name>Joe Smith</name><email>xxx@sina.com</email></User>%
	})

	e.Logger.Fatal(e.Start(":1323"))
}
