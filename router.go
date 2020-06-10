package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Init is main router
func Init() *echo.Echo {

	e := echo.New()

	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signup", handler.Signup) // POST signup
	e.POST("/login", handler.Login)   // POST login

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/ranking", handler.Getdata)               // GET api/ranking
	api.POST("/addpost", handler.Addpost)              // POST api/addpost 勉強した分加えるやつ
	api.DELETE("/addpost/:id", handler.Deletepost)     // DELETE api/:id 消す
	api.PUT("/addpost/:id/config", handler.Updatepost) // PUT addpost/:id/config 変更する

	return e
}
