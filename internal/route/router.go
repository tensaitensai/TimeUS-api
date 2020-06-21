package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensaitensai/TimeUS-api/internal/handler"
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

	//e.GET("/ranking", handler.GetRanking) // GET ranking ランキング
	//e.GET("/user/:uid/post", handler.GetPost) // GET user/:uid/post その人の勉強分をもらうやつ

	api := e.Group("/api") //
	api.Use(middleware.JWTWithConfig(handler.Config))
	//api.PUT("/api/user/:uid") // PUT api/user/:uid　プロフィール変更でつかう
	api.POST("/user/:uid/post", handler.AddPost)          // POST api/user/:uid/post 勉強した分加えるやつ
	api.DELETE("/user/:uid/post/:id", handler.DeletePost) // DELETE api/user/:uid/post/:id post消す
	api.PUT("/user/:uid/post/:id", handler.UpdatePost)    // PUT /user/:uid/post/:id 変更する

	return e
}
