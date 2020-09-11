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

	e.GET("/user/:uid/post", handler.GetPost) // GET user/:uid/post その人の勉強分をもらうやつ

	//e.GET("/ranking", handler.GetRanking) // GET ranking ランキング

	api := e.Group("/api") //
	api.Use(middleware.JWTWithConfig(handler.Config))
	//api.PUT("user/")                            // PUT api/user/　プロフィール変更でつかう
	api.POST("/post", handler.AddPost)          // POST api/post 勉強した分加えるやつ
	api.DELETE("/post/:id", handler.DeletePost) // DELETE api/post/:id post消す
	api.PUT("/post/:id", handler.UpdatePost)    // PUT /user/:uid/post/:id 変更する

	return e
}
