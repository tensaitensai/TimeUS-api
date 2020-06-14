package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensaitensai/TimeUS-api/handler"
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

	/*
		api := e.Group("/api")
		api.Use(middleware.JWTWithConfig(handler.Config))
		api.GET("/ranking", handler.Getdata)               // GET api/ranking
		api.POST("/post", handler.Addpost)              // POST api/post 勉強した分加えるやつ
		api.DELETE("/post/:id", handler.Deletepost)     // DELETE api/post/:id 消す
		api.PUT("/apost/:id", handler.Updatepost) // PUT api/post/:id 変更する
	*/
	return e
}
